import time
from dotenv import load_dotenv

from .controller.grpc.router import init_routes
from .pkg.config import get_config
from .pkg.logger import get_logger
from .server.grpc.config import GRPCConfig
from .server.grpc.server import GRPCServer
from .service.service import Service

# Get config and logger
load_dotenv(".env")
conf = get_config()
logger = get_logger()


def create_app():
    # Init services
    service = Service()

    # Init routes
    routes = init_routes(service)

    # Init grpc server
    server = GRPCServer(GRPCConfig(
        host=conf.scrapper_host,
        port=conf.scrapper_port,
        count_workers=conf.scrapper_count_workers
    ), routes)

    # Start grpc server
    try:
        server.start()
        logger.info(f"server started on port {conf.scrapper_port}")
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(0)
    except Exception as e:
        logger.error(f'Error while starting grpc server: {e}')
        server.stop(0)
