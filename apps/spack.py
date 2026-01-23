if [ -f ${HOME}/spack/share/spack/setup-env.sh ]; then
    . ${HOME}/spack/share/spack/setup-env.sh
    echo "\033[32m\033[0;39m  spack activated"
else
    echo "\033[31m\033[0;39m  spack not found"
fi