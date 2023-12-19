from controller.grpc.file import FileService


def init_routes():
    return [FileService()]
