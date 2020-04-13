from flask import Flask,request
app = Flask(__name__)
@app.route('/20161021/Accounts/35971af03f12b9ff8d2d47dd48a932f6/AccountInfo?sig=A4BCE165F30DE94CFBE1F794A9FF3746', methods=['POST'])
def register():
	print(request.form.get('name'))
	print(request.form.get('age'))
	
	json_dict = {
		response:{
			"retcode":"0000",
			"number":"1002"
		}
	}

	return 

if __name__ == '__main__':
	app.run()
