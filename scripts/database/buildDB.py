import pyaes, base64, os, yaml, pymongo, subprocess

def getDecryptedValue(key, value):
	decoded = base64.b64decode(value);
	iv = decoded[:16]
	cypherText = decoded[16:]
	decrypter = pyaes.Decrypter(pyaes.AESModeOfOperationCBC(key, iv=iv))
	return (decrypter.feed(cypherText) + decrypter.feed()).decode('utf8')

def getConfig():
	key = None
	config = None
	with open('config.yaml.key', 'rb') as file:
		key=file.read()
		file.close()
	with open('config.yaml', 'r') as file:
		config = yaml.safe_load(file)
		file.close()
	config['password'] = getDecryptedValue(key, config['password'])
	createUser(config)

def createUser(config):
	client = pymongo.MongoClient(config['server'], config['port'])
	client.admin.authenticate('admin', 'defaultPassword')
	client.admin.add_user('admin', config['adminPassword'])
	client.admin.logout()
	client.admin.authenticate('admin', config['adminPassword'])
	db = client[config['database']]
	db.add_user(config['username'], config['password'], roles=[{'role':'readWrite','db':config['database']}])
	db.logout()

getConfig()
