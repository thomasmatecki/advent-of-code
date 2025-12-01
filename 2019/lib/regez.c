#include "regez.h"
#include <regex.h>
#include <stdlib.h>
#include <stdalign.h>

// Internal wrapper struct that contains the actual regex_t
struct regex_wrapper
{
    regex_t regex;
};

regex_t_ptr regex_compile(const char *pattern, int cflags)
{
    struct regex_wrapper *wrapper = malloc(sizeof(struct regex_wrapper));
    if (!wrapper)
    {
        return NULL;
    }

    int ret = regcomp(&wrapper->regex, pattern, cflags);
    if (ret != 0)
    {
        free(wrapper);
        return NULL;
    }

    return wrapper;
}

int regex_exec(regex_t_ptr regex, const char *string, size_t nmatch, void *pmatch, int eflags)
{
    if (!regex)
    {
        return -1;
    }
    return regexec(&regex->regex, string, nmatch, (regmatch_t *)pmatch, eflags);
}

void regex_free(regex_t_ptr regex)
{
    if (regex)
    {
        regfree(&regex->regex);
        free(regex);
    }
}

size_t regex_sizeof(void)
{
    return sizeof(regex_t);
}

size_t regex_alignof(void)
{
    return alignof(regex_t);
}