from webscrapper_proto.gen.python.scrapper.scrapper_pb import *
from webscrapper_proto.gen.python.scrapper.scrapper_pb_grpc import *


class FileService(FileServiceServicer):
    def GetFile(self, request, context):
        # Здесь вы можете реализовать логику получения файла
        with open(request.name, 'rb') as f:
            content = f.read()
        return FileResponse(content=content)
