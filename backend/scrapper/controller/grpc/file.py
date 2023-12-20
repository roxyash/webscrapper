import uuid

from gen.python.scrapper.scrapper_pb import FileResponse
from gen.python.scrapper.scrapper_pb_grpc import FileServiceServicer


class FileHandler(FileServiceServicer):
    def __init__(self, file):
        self._file = file

    def GetFile(self, request, context):
        """
        Get file by params
        :param request:
        :param context:
        :return:
        """
        # Получение файла
        file_content = self._file.scrap('2023-12-19', '2023-12-19', 'H.danil007@yandex.ru', '100300qW',
                                        str(uuid.uuid4()))

        # Создание и возврат FileResponse
        return FileResponse(content=file_content)
