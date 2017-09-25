# Tools


This is a set of tools and scripts for `go-coding` project.


<a name="docker"><br /></a>
## Docker Notes

### Dockerfile ENTRYPOINT vs CMD

  - No `ENTRYPOINT`

    | CMD form             | Actual calling       |
    |:---------------------|:---------------------|
    | No `CMD`             | *error, not allowed* |
    | `CMD cmd arg`        | /bin/sh -c cmd arg   |
    | `CMD ["cmd", "arg"]` | cmd arg              |

  - Shell form `ENTRYPOINT exec param`

    | CMD form             | Actual calling                           |
    |:---------------------|:-----------------------------------------|
    | No `CMD`             | /bin/sh -c exec param                    |
    | `CMD cmd arg`        | /bin/sh -c exec param /bin/sh -c cmd arg |
    | `CMD ["cmd", "arg"]` | /bin/sh -c exec param cmd arg            |

  - Exec form: `ENTRYPOINT ["exec", "param"]`

    | CMD form             | Actual calling                |
    |:---------------------|:------------------------------|
    | No `CMD`             | exec param                    |
    | `CMD cmd arg`        | exec param /bin/sh -c cmd arg |
    | `CMD ["cmd", "arg"]` | exec param cmd arg            |
