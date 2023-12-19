from concurrent import futures

import grpc

from config import Config


class GRPCServer:
    def __init__(self, config: Config, handlers):
        self.server = self._init_server(handlers)
        self.config = config

    def start(self):
        """
        Method for starting grpc server
        """
        self.server.start()

    def stop(self, grace: int = 0):
        """
        Method for stopping grpc server
        :param grace: value in seconds
        :return:
        """
        self.server.stop(grace)

    def wait_for_termination(self, timeout: int = 0):
        """
        Method for waiting for termination grpc server
        :param timeout: value in seconds
        :return:
        """
        self.server.wait_for_termination(timeout)

    def _init_server(self, handlers: str):
        """
        Method for initializing grpc server
        :param handlers: list of handlers
        :return: grpc server
        """
        server = grpc.server(futures.ThreadPoolExecutor(max_workers=self.config.count_workers))
        for handler in handlers:
            print(handler)
            # file_pb2_grpc.add_FileServiceServicer_to_server(handler, server)
        server.add_insecure_port(f'[::]:{self.config.port}')
        return server
