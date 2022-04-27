# IPv6 Update Monitor and Notifier


## Abstract

It's currently still in prototype state but works as expected. I wrote this up for my raspberry pi which has a changing ip address.
This daemon is listening for netlink events by using the `ip` command. If it gets notified about an ip change at the given device, it
passes this new ip to an external command of your choice. In my setup this command simply executes via SSH a remote `nsupdate`.

It is easy to set up an own dyndns service with that.

Currently it needs still a shell wrapper. The functionality from the wrapper should be implemented directly in the daemon.

This software is still a proof of concept and just a prototype, but feel free to fork and make it final ;).


## License

MIT
