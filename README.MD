# Searching fuzzing targets for Golang code

A utility for searching and prioritizing targets for fuzzing. The level of granularity is a function.

## Implemented Algorithms

### Metrics based

On a project, a metric is calculated for each function - score of function. then the functions are ranked based on it and their priority is calculated.

There are merics that implemented:

+ Complexity Metrics
    + Function
        + Cyclomatic complexity
    + Loop Structures
        + Number of loops
        + Loop Structures C3 #ofnested loops
        + Maximumnesting level of loops
+ Vulnerability Metrics
    + Dependency
        + Number ofparameter variables
        + Number ofvariables as parameters for callee function
    + Pointers
        + Number ofpointer arithmetic
        + Number of variables involved in pointer arithmetic
        + Maxpointer arithmetic a variable is involved in
    + Control Structures
        + Number of nested control structures
        + Maximum nesting level of control structures
        + Maximum of control-dependent control structures
        + Maximum of data-dependent control structures
        + Number of if structures without else
        + Number of variables involved in control predicates

### LLM based

TODO

## How to use

It is console application with this arguments:
+ folder - path to the project
+ package - name of analyzing package (Default: none <=> all packages)
+ algorithm - name of algorithm (complexity/volnerability) (Default: volnerability)
+ top - count of output function with max priority (Default: 5)

## Examples of using

### https://github.com/kataras/go-serializer

#### Input
```
--folder={workingDir}/go-serializer
--top=5
```

#### Output

```

```

### https://github.com/google/syzkaller

#### Input

```
--folder={workingDir}/syzkaller
--top=8
--algorithm=complexity
--package=email
```

#### Output

```
Package: github.com/google/syzkaller/pkg/email, Name: Parse, Priority: 100
Package: github.com/google/syzkaller/pkg/email, Name: extractCommand, Priority: 95
Package: github.com/google/syzkaller/pkg/email, Name: ParsePatch, Priority: 90
Package: github.com/google/syzkaller/pkg/email, Name: parseBody, Priority: 81
Package: github.com/google/syzkaller/pkg/email, Name: extractBodyBugIDs, Priority: 76
Package: github.com/google/syzkaller/pkg/email, Name: strToCmd, Priority: 71
Package: github.com/google/syzkaller/pkg/email, Name: MergeEmailLists, Priority: 67
Package: github.com/google/syzkaller/pkg/email, Name: extractArgsTokens, Priority: 62
```

### https://github.com/googleapis/google-api-go-client

#### Input

```
--folder={workingDir}/google-api-go-client
--top=10
--algorithm=volnerability
```

#### Output

```
Package: google.golang.org/api/google-api-go-generator, Name: generateCode, Priority: 100
Package: google.golang.org/api/internal/gensupport, Name: schemaToMap, Priority: 100
Package: google.golang.org/api/transport/bytestream/internal, Name: Write, Priority: 100
Package: google.golang.org/api/transport/grpc, Name: dial, Priority: 100
Package: google.golang.org/api/internal/gensupport, Name: Upload, Priority: 100
Package: google.golang.org/api/internal/gensupport, Name: sendAndRetry, Priority: 100
Package: google.golang.org/api/google-api-go-generator, Name: GenerateCode, Priority: 100
Package: google.golang.org/api/examples, Name: gmailMain, Priority: 100
Package: google.golang.org/api/impersonate, Name: CredentialsTokenSource, Priority: 100
Package: google.golang.org/api/examples, Name: fitnessMain, Priority: 100
```