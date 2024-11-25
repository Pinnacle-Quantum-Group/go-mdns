
#include "mdns_wrapper.h"

// Wrapper for starting mDNS service
void start_mdns_service(const char *service_name, const char *hostname, int port) {
    mdns_start_service(service_name, hostname, port);
}

// Wrapper for stopping mDNS service
void stop_mdns_service() {
    mdns_stop_service();
}
