from concurrent import futures

from gen.python.scrapper.scrapper_pb_grpc import *

from .config import GRPCConfig


class GRPCServer:
    def __init__(self, config: GRPCConfig, handlers):
        self._config = config
        self._server = self._init_server(handlers)

    def start(self):
        """
        Method for starting grpc server
        """
        self._server.start()

    def stop(self, grace: int = 0):
        """
        Method for stopping grpc server
        :param grace: value in seconds
        :return:
        """
        self._server.stop(grace)

    def wait_for_termination(self, timeout: int = 0):
        """
        Method for waiting for termination grpc server
        :param timeout: value in seconds
        :return:
        """
        self._server.wait_for_termination(timeout)

    def _init_server(self, handlers: str):
        """
        Method for initializing grpc server
        :param handlers: list of handlers
        :return: grpc server
        """
        server = grpc.server(futures.ThreadPoolExecutor(max_workers=self._config.count_workers))
        for handler in handlers:
            add_FileServiceServicer_to_server(handler, server)
        server.add_insecure_port(f'[::]:{self._config.port}')
        return server
