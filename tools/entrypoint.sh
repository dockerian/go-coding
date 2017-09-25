#!/bin/bash

# Add local user and group

USER=${USER:-user}
USER_ID=${LOCAL_USER_ID:-9001}
GROUP_ID=${LOCAL_GROUP_ID:-9001}

echo "Starting entrypoint with user ${USER} [UID= $USER_ID, GID= ${GROUP_ID}]"

if [[ "${USER_ID}" == "0" ]]; then
  USER="root"
  export HOME=/root
else
  groupadd --gid ${GROUP_ID} -o ${USER}
  useradd --shell /bin/bash -u ${USER_ID} -o -c "container user ${USER}" -m ${USER} --gid ${GROUP_ID}
  echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
  export HOME=/home/${USER}
  echo "export PS1='\n\u@\h [\w] \D{%F %T} [\#]:\n\$ '" >> ${HOME}/.bashrc
  echo "alias ll='ls -al'" >> ${HOME}/.bashrc
  echo "" >> ${HOME}/.bashrc
fi

exec /usr/local/bin/gosu ${USER} "$@"
