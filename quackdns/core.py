# Library of functions interacting with DuckDNS
from abc import ABC, abstractmethod
import requests
import time

__DUCK_DNS_URL__ = 'https://www.duckdns.org/update'


class AbstractUpdater(ABC):
    """
    Abstract class that defines the methods
    every Updater should have.
    """

    def __init__(self,
                 domains: str,
                 token: str,
                 ip: str=None,
                 ipv6: str=None,
                 verbose='true',
                 clear=None):
        """

        :param domains: required - comma separated
        list of the subnames you want to update
        :param token: required - your account token
        :param ip: optional - an ipv4 address,
        if left blank the ip is detected automatically
        :param ipv6: optional - an ipv6 address,
        if left blank the ip is detected automatically
        :param verbose: optional - defaults to 'true',
        else it can be set to 'false'
        :param clear: optional - if set to 'true'it clears
        any data about the ip addresses
        """
        super().__init__()

        # Initialise parameters
        self.set_domains(domains)
        self.__token__ = token
        self.__ip__ = ip
        self.__ipv6__ = ipv6
        self.__verbose__ = verbose
        self.__clear__ = clear

        self.__dns_url__ = __DUCK_DNS_URL__
        self.__verify_secure_connection__ = True

    def is_valid_domains(self, domains):
        if domains is None or domains == "":
            return False

        domains_list = domains.split(",")

        for domain in domains_list:
            # todo: use a regex
            if domain == "":  # Empty string
                return False
            if " " in domain:  # Contains whitespaces
                return False

        return True

    def set_domains(self, domains):
        if self.is_valid_domains(domains):
            self.__domains__ = domains
        else:
            raise InvalidParameterError("domains")

    def get_dns_url(self):
        return self.__dns_url__

    def get_verify_secure_connection(self):
        return self.__verify_secure_connection__

    def get_params_dict(self):
        params = dict()
        params['domains'] = self.__domains__
        params['token'] = self.__token__

        if self.__ip__:
            params['ip'] = self.__ip__

        if self.__ipv6__:
            params['ipv6'] = self.__ipv6__

        if self.__verbose__:
            params['verbose'] = self.__verbose__

        if self.__clear__:
            params['clear'] = self.__clear__

        return params

    def loop_update(self, sleep_seconds):
        while True:
            self.update()
            time.sleep(sleep_seconds)

    @abstractmethod
    def update(self):
        pass


class MockUpdater(AbstractUpdater):
    """
    Mock updater implementation which fakes a class
    to update the DuckDNS information.
    """

    def __init__(self):
        super().__init__("mock-domain", "mock-token")

    def update(self):
        response = "Updated!"
        return response


class Updater(AbstractUpdater):
    """
    Default implementation of an Updater.

    It sends an HTTP GET request to the default server.
    """

    def __init__(self,
                 domains,
                 token,
                 ip=None,
                 ipv6=None,
                 verbose='true',
                 clear=None):
        super().__init__(domains,
                         token,
                         ip=None,
                         ipv6=None,
                         verbose='true',
                         clear=None)
        pass

    def update(self):
        params = self.get_params_dict()
        response = requests.get(self.get_dns_url(),
                                params=params,
                                verify=self.get_verify_secure_connection())
        return response.text


class InvalidParameterError(Exception):
    def __init__(self, invalid_parameter, reason=None):
        self.msg = "Parameter '{}' is not valid.".format(invalid_parameter)
        if reason:
            self.msg += "Reason: {}.".format(reason)

    def __str__(self):
        return repr(self.msg)
