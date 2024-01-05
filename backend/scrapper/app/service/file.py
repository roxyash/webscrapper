import logging
import os
import shutil
import time

from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from seleniumwire import webdriver

logging.getLogger('seleniumwire').setLevel(logging.ERROR)


class FileService:
    def __init__(self):
        ...

    @staticmethod
    def scrap(date_from: str, date_to: str, login: str, password: str, user_id: str):
        """
        Функция для скачивания отчета "Сводка по товарам" с mpassistant.pro
        :param login: Your mpassistant.pro login
        :param password: Your mpassistant.pro password
        :param date_from: Format: dd.mm.yyyy
        :param date_to: Format dd.mm.yyyy
        :param user_id: Unique user identifier
        :return:
        """

        # Создаем объект опций и добавляем аргумент для режима "headless"
        options = Options()
        options.add_argument('--headless')

        # Указываем директорию для скачивания файлов
        download_dir = os.path.join(os.getcwd(), 'files', user_id)
        os.makedirs(download_dir, exist_ok=True)
        prefs = {'download.default_directory': download_dir}
        options.add_experimental_option('prefs', prefs)

        # Инициализируем драйвер с созданными опциями
        driver = webdriver.Chrome(seleniumwire_options={'enable_har': True}, options=options)

        try:
            # Навигация к странице логина
            driver.get('https://app.mpassistant.pro/app/login')

            # Заполнение полей формы логина и отправка формы
            driver.find_element(By.ID, 'login').send_keys(login)
            driver.find_element(By.ID, 'password').send_keys(password)
            driver.find_element(By.XPATH, "//button[@type='submit']").click()

            # Ожидание авторизации
            time.sleep(5)
            # Переход на страницу отчета
            driver.get(
                f'https://app.mpassistant.pro/app/report/product?filter%5Bstart_period%5D={date_from}&filter%5Bend_period%5D={date_to}&filter%5Barticle%5D=&filter%5Bmarketplace%5D=&filter%5Bshop%5D=&filter%5Bbrand%5D=&filter%5BsortBy%5D=profit&filter%5BorderBy%5D=desc&filter%5Bview%5D=product&filter%5Bself_ransom%5D=include&filter%5Barchive_product%5D=exclude&filter%5Bactivity%5D=yes')

            # Отправка запроса на скачивание файла
            download_button = driver.find_element(By.CLASS_NAME, 'ajax_export_excel.btn.btn-outline-success')
            driver.execute_script("arguments[0].click();", download_button)

            # Ожидание скачивания файла
            time.sleep(5)

            # Чтение файла и возвращение его содержимого
            file_path = os.path.join(download_dir, 'export.xls')
            with open(file_path, 'rb') as f:
                file_bytes = f.read()

            # Удаление файла после чтения
            shutil.rmtree(download_dir)

            return file_bytes

        finally:
            # Закрываем браузер независимо от успешности скачивания файла
            driver.quit()
