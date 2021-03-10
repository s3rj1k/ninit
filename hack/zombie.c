#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>

/*
  https://vitux.com/how-to-create-a-dummy-zombie-process-in-ubuntu/

  compile staticaly: `cc -static zombie.c -o zombie`
  start binary and check for zombies: `ps axo stat,ppid,pid,comm | grep -w defunct`
*/

int main ()
{
  pid_t child_pid;child_pid = fork ();
  if (child_pid > 0) {
    sleep (60);
  } else {
    exit (0);
  }
  return 0;
}
