#ifndef MDNS_WRAPPER_H
#define MDNS_WRAPPER_H

#include "mdns.h"

// Wrapper functions for Go compatibility
void start_mdns_service(const char *service_name, const char *hostname, int port);
void stop_mdns_service();

#endif

