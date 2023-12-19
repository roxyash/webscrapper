import time

from dotenv import load_dotenv
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By

from pkg.config import get_config
from pkg.logger import get_logger

# Get config and logger
load_dotenv(".env")
conf = get_config()
logger = get_logger()


def download_report(date_from: str, date_to: str) -> None:
    """
    Функция для скачивания отчета "Сводка по товарам" с mpassistant.pro
    :param login: Your mpassistant.pro login
    :param password: Your mpassistant.pro password
    :param date_from: Format: dd.mm.yyyy
    :param date_to: Format dd.mm.yyyy
    :return:
    """

    # Создаем объект опций и добавляем аргумент для режима "headless"
    options = Options()
    options.add_argument('--headless')

    # Инициализируем драйвер с созданными опциями
    driver = webdriver.Chrome(options=options)

    try:
        # Навигация к странице логина
        driver.get('https://app.mpassistant.pro/app/login')

        # Заполнение полей формы логина и отправка формы
        driver.find_element(By.ID, 'login').send_keys(conf.mpassistant_user)
        driver.find_element(By.ID, 'password').send_keys(conf.mpassistant_password)
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
    finally:
        # Закрываем браузер независимо от успешности скачивания файла
        driver.quit()


def main():
    print(conf)


if __name__ == '__main__':
    main()
