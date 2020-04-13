import json
import re
import sys
from pathlib import Path

DIRECTIVE_RE = re.compile('\| cc_pc(.*)\|')

def batch_call_log(match, infile, outfile):
    first_line = 1
    tmp_customer = 0
    lost_num = 0
    result = match.groups()[0]
    print(result)
    customer = int(result.split('|')[1])
    if first_line:
        tmp_customer = customer
        first_line = 0
    else:
        while tmp_customer != customer:
            tmp_customer += 1
            if(tmp_customer != customer):
                lost_num += 1
                outfile.write(str(tmp_customer)+'\n')
    outfile.write("total num:" + str(lost_num))

class TemplateEngine:
    def __init__(self, infilename, outfilename, log_handle):
        self.infile = infilename
        self.pos = 0
        self.outfile = outfilename
        self.log_handle = log_handle
    def process(self):
        infile = open(self.infile)
        outfile = open(self.outfile, 'w')
        template = infile.read()

        match = DIRECTIVE_RE.search(template, pos=self.pos)
        while match:
            self.log_handle(match, infile, outfile)
            self.pos = match.end()
            match = DIRECTIVE_RE.search(template, pos=self.pos)
        outfile.close()
        infile.close()
        
        
e = TemplateEngine("infile", "outfile", batch_call_log)
e.process()
