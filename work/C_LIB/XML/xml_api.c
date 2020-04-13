#include <libxml/xpath.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

static int XMLDocParser(char *fileName){
	int ret = 0;
	xmlDocPtr doc = NULL;
	xmlNodePtr node = NULL;

	xmlKeepBlanksDefault(0);
	//libxml默认将各个节点间的空格当做一个节点，只要在调用xmlParse之前调用xmlParseFile之前调用
	//xmlKeepBlanksDefault(0) 就可以了

	doc = xmlParseFile(fileName);
	if(doc == NULL){
		ret = -1;
		return ret;
	}

	node = xmlDocGetRootElement(doc);
	if(node == NULL){
		ret = -1;
		return ret;
	}

	if(xmlStrcmp(node->name, (const xmlChar *)"root")){
		ret = -1;
		return ret;
	}

	printf("node->line: %d\n", node->line);
	
	node = node->children;
	while(node != NULL){
		printf("node->line: %d\n", node->line);
		printf("name=%s content=%s\n", node->name, (char *)xmlNodeGetContent(node));
		node = node->next;
	}


_exit:
	if(doc){
		xmlFreeDoc(doc);
	}
	return ret;
}

int main(){
	char *docname = "story.xml";
	XMLDocParser(docname);
	return 0;
}
