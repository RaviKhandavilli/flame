import json
import sys

CONF_KEY_MQTT_BROKER = 'broker'
CONF_KEY_BACKEND = 'backend'
CONF_KEY_AGENT = 'agent'
CONF_KEY_JOB = 'job'
CONF_KEY_JOB_ID = 'id'
CONF_KEY_JOB_NAME = 'name'
CONF_KEY_ROLE = 'role'
CONF_KEY_REALM = 'realm'
CONF_KEY_CHANNEL = 'channel'

CONF_KEY_CHANNEL_NAME = 'name'
CONF_KEY_CHANNEL_PAIR = 'pair'
CONF_KEY_CHANNEL_IS_BIDIR = 'isBidirectional'
CONF_KEY_CHANNEL_GROUPBY = 'groupBy'
CONF_KEY_CHANNEL_GROUPBY_TYPE = 'type'
CONF_KEY_CHANNEL_GROUPBY_VALUE = 'value'

GROUPBY_DEFAULT_GROUP = 'default'

BACKEND_TYPE_LOCAL = 'local'
BACKEND_TYPE_P2P = 'p2p'
BACKEND_TYPE_MQTT = 'mqtt'

REALM_SEPARATOR = '.'

backend_types = [BACKEND_TYPE_LOCAL, BACKEND_TYPE_P2P, BACKEND_TYPE_MQTT]


class Config(object):
    class Job(object):
        def __init__(self, json_data=None):
            self.job_id = json_data[CONF_KEY_JOB_ID]
            self.name = json_data[CONF_KEY_JOB_NAME]

        def print(self):
            print('\t--- job ---')
            print(f'\t{CONF_KEY_JOB_ID}: {self.job_id}')
            print(f'\t{CONF_KEY_JOB_NAME}: {self.name}')

    class Channel(object):
        class GroupBy(object):
            def __init__(self, json_data=None):
                self.gtype = ''
                self.value = []

                if json_data is None:
                    return

                self.gtype = json_data[CONF_KEY_CHANNEL_GROUPBY_TYPE]
                self.value = json_data[CONF_KEY_CHANNEL_GROUPBY_VALUE]

            def groupable_value(self, realm=''):
                for entry in self.value:
                    # check if an entry is a prefix of realm in a dot-separated
                    # fashion; if so, then return the matching entry
                    if realm.startswith(entry) and (
                        len(realm) == len(entry) or
                        realm[len(entry)] == REALM_SEPARATOR
                    ):
                        return entry

                return GROUPBY_DEFAULT_GROUP

            def print(self):
                print('\t\t--- groupby ---')
                print(f'\t\t{CONF_KEY_CHANNEL_GROUPBY_TYPE}: {self.gtype}')
                print(f'\t\t{CONF_KEY_CHANNEL_GROUPBY_VALUE}: {self.value}')

        def __init__(self, json_data):
            self.name = json_data[CONF_KEY_CHANNEL_NAME]
            self.pair = json_data[CONF_KEY_CHANNEL_PAIR]
            if len(self.pair) != 2:
                sys.exit(f'A pair must have two ends, but got {self.pair}')

            self.is_bidrectional = True
            if CONF_KEY_CHANNEL_IS_BIDIR in json_data:
                self.is_bidrectional = json_data[CONF_KEY_CHANNEL_IS_BIDIR]

            if CONF_KEY_CHANNEL_GROUPBY in json_data:
                self.groupby = Config.Channel.GroupBy(
                    json_data[CONF_KEY_CHANNEL_GROUPBY]
                )
            else:
                self.groupby = Config.Channel.GroupBy()

        def print(self):
            print('\t--- channel ---')
            print(f'\t{CONF_KEY_CHANNEL_NAME}: {self.name}')
            print(f'\t{CONF_KEY_CHANNEL_PAIR}: {self.pair}')
            print(f'\t{CONF_KEY_CHANNEL_IS_BIDIR}: {self.is_bidrectional}')
            self.groupby.print()

    def __init__(self, config_file: str):
        with open(config_file) as f:
            json_data = json.load(f)
            f.close()

        self.broker = json_data[CONF_KEY_MQTT_BROKER]
        self.backend = json_data[CONF_KEY_BACKEND]
        if not self.is_valid(self.backend, backend_types):
            sys.exit(f'not a vailid backend type: {self.backend}')

        self.agent = json_data[CONF_KEY_AGENT]
        self.job = Config.Job(json_data[CONF_KEY_JOB])
        self.role = json_data[CONF_KEY_ROLE]
        self.realm = json_data[CONF_KEY_REALM]
        self.channels = {}
        for channel_info in json_data[CONF_KEY_CHANNEL]:
            channel_config = Config.Channel(channel_info)
            self.channels[channel_config.name] = channel_config

    def is_valid(self, needle, haystack):
        return needle in haystack

    def print(self):
        print('--- config ---')
        print(f'{CONF_KEY_MQTT_BROKER}: {self.broker}')
        print(f'{CONF_KEY_BACKEND}: {self.backend}')
        print(f'{CONF_KEY_AGENT}: {self.agent}')
        print(f'{CONF_KEY_ROLE}: {self.role}')
        print(f'{CONF_KEY_REALM}: {self.realm}')
        self.job.print()
        for _, channel in self.channels.items():
            channel.print()
