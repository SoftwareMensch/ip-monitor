#!/bin/sh

# some invalid lines
echo -e -n "Deleted aaaa:aaaa:aaaa:aaa::/64 proto ra metric 1024 expires 7191sec pref medium \n"
echo -e -n "Deleted 2: eth0    inet6 aaaa:aaaa:aaaa:aaa:bbbb:bbbb:bbbb:bbbb/64 scope global deprecated dynamic mngtmpaddr noprefixroute \n"
echo -e -n "valid_lft 7191sec preferred_lft 0sec \n"
echo -e -n "2: eth0    inet6 fd00::aaaa:aaaa:aaaa:aaaa/64 scope global deprecated dynamic mngtmpaddr noprefixroute \n"
echo -e -n "2: eth0    inet6 fd00::aaaa:aaaa:aaaa:aaaa/64 scope global tentative dynamic mngtmpaddr noprefixroute \n"
echo -e -n "what ever \n"

# valid line
echo -e -n "2: eth0    inet6 aaaa:aaaa:aaaa:aaaa:bbbb:bbbb:bbbb:bbbb/64 scope global tentative dynamic mngtmpaddr noprefixroute \n"

# some more invalid lines
echo -e -n "valid_lft 7191sec preferred_lft 0sec \n"
echo -e -n "2: eth0    inet6 fd00::aaaa:aaaa:aaaa:aaaa/64 scope global deprecated dynamic mngtmpaddr noprefixroute \n"

# EOF
