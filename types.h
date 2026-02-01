#ifndef TYPES_H
#define TYPES_H

typedef struct {
    char* data;
    int len;
} PascalString;

typedef struct {
    PascalString* data;
    int len;
} PascalStringArray;

#endif
