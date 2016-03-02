package resolve

import (
	"syscall"
    "net"
    "fmt"
)

type stringError struct {
    what string
}

func (e *stringError) Error() string {
    return e.what
}

func newError(text string) error {
    return &stringError{text}
}

func findIP(ips []net.IP, af int) (ip net.IP, e error) {
    var theIP net.IP
    if af == syscall.AF_INET {
        // AF_INET
        for _, ip := range ips {
            v := ip.To4() 
            if v != nil {
                theIP = v
                break
            }
        }
        
        if theIP == nil {
            return nil, newError("No IPv4 available for this host")
        }
        return theIP, nil
    }
    
    // AF_INET6
    for _, ip := range ips {
        v := ip.To4() 
        if v == nil {
            theIP = ip.To16()
            break
        }
    }
    
    if theIP == nil {
        return nil, newError("No IPv6 available for this host")
    }
    return theIP, nil
}

// GetAddrInfo is used as a substitute for the glibc getaddrinfo when using syscall.Connect
func GetAddrInfo(host string, port, af int) (sa syscall.Sockaddr, e error) {
    
    if (af != syscall.AF_INET) && (af != syscall.AF_INET6) {
        return nil, newError(fmt.Sprintf("Address family %v is not supported!", af))
    }
    
    ips, e := net.LookupIP(host)
    
    if e != nil {
        return
    }
    
    ip, e := findIP(ips, af)
    
    if e != nil {
        return 
    }
    
    if af == syscall.AF_INET {
        sa4 := &syscall.SockaddrInet4{Port: port}
        copy(sa4.Addr[:], ip[0:4])
        return sa4, nil
    }
    sa6 := &syscall.SockaddrInet6{Port: port}
    copy(sa6.Addr[:], ip[0:16])
    return sa6, nil
}
