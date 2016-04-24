#include <stdio.h>
#import <stdio.h>
#import <unistd.h>
#import <sys/socket.h>
#import <sys/sys_domain.h>
#import <net/if_utun.h>
#import <sys/kern_control.h>
#import <sys/ioctl.h>
#import <string.h>
#import <sys/errno.h>

void osxtun_open(int *fd, int *unit, char **err)
{
    // open socket
    *fd = socket(PF_SYSTEM, SOCK_DGRAM, SYSPROTO_CONTROL);
    
    struct ctl_info ci;
    snprintf(ci.ctl_name, sizeof(ci.ctl_name), UTUN_CONTROL_NAME);
    if (ioctl(*fd, CTLIOCGINFO, &ci) == -1) {
        *err = strerror(errno);
        close(*fd);
        *fd = -1;
    }
    
    struct sockaddr_ctl sc;
    sc.sc_id = ci.ctl_id;
    sc.sc_len = sizeof(sc);
    sc.sc_family = AF_SYSTEM;
    sc.ss_sysaddr = AF_SYS_CONTROL;
    
    int unit_nr = 0;
    do {
        sc.sc_unit = unit_nr + 1;
        if (!connect(*fd, (struct sockaddr * )&sc, sizeof(sc))) {
            break;
        }
        unit_nr++;
        
    } while (sc.sc_unit < 255);
    
    if (unit_nr > 254) {
        *err = strerror(errno);
        close(*fd);
        *fd = -1;
        return;
    }

    *unit = unit_nr;
}
