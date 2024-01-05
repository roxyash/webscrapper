class GRPCConfig:
    def __init__(self, host: str, port: str, count_workers: int):
        self.host = host
        self.port = port
        self.count_workers = count_workers
