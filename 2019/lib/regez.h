#ifndef REGEZ_H
#define REGEZ_H

#include <stddef.h>
#include <regex.h>

// Opaque pointer type for regex - hides the actual regex_t from Zig
typedef struct regex_wrapper *regex_t_ptr;

regex_t_ptr regex_compile(const char *pattern, int cflags);

// Allocate and compile a regex pattern
// Returns NULL on error
// Execute regex match
// Returns 0 on match, non-zero otherwise
int regex_exec(regex_t_ptr regex, const char *string, size_t nmatch, void *pmatch, int eflags);

// Free regex resources
void regex_free(regex_t_ptr regex);

// Get size information (useful for manual memory management)
size_t regex_sizeof(void);
size_t regex_alignof(void);

#endif // REGEZ_H