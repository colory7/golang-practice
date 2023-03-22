#include <stdio.h>
#include <string.h>

int demo(char * s)
{
	int len = strlen(s);
    int map[256] = {0};				//哈希表
    int res = 0, l = 0, r = 0;
	while(r < len){					//遍历字符串
		if(map[s[r]] == 0){			//不是重复字符，右指针向右移动
            map[s[r++]]++;
        }else{						//遇到重复字符，左指针向左移动
			map[s[l++]]--;
        }
        res = len;		//更新最大长度
    }
	 char a = 'a';
	 s = &a;
     printf("s: %s\n",s);

    return res;
}