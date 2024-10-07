#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void swap(char a, char b) {
    char tmp = a;
    a =b;
    b = tmp;
}
char * reverse(char *s) {
   // char *rev= (char *)malloc(sizeof(char) *strlen(s));
    int i=0;
    int len = strlen(s);
    while(i<len/2) {
        char tmp = s[i];
        s[i] = s[len-i-1];
        s[len-1-i] = tmp;
        i++;
    }
    return s;
}
int main() {
    char str[] = "abc";
    char *rev= reverse(str);
    printf("%s\n", rev);
    return 0;
}