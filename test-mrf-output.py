#!/usr/bin/python

import json
import sys

def main(argv=sys.argv):
	if len(argv) < 2:
		print("nope, need a filename")
		sys.exit(1)

	with open(argv[1]) as log_file:
		recorded = {}
		one_json_line = log_file.readline()
		while one_json_line:
			print("read")
			one_list = json.loads(one_json_line.strip())
	
			for one_obj in one_list:
				if one_obj['namespace'] in recorded and recorded[one_obj['namespace']] == one_obj['data']:
					print("Filtering not done, duplicate metric: {namespace} value: {value}".format(namespace=one_obj['namespace'], value=one_obj['data']))
					sys.exit(0)
				else:
					recorded[one_obj['namespace']] = one_obj['data']
			one_json_line = log_file.readline()

if __name__=="__main__":
	main()
