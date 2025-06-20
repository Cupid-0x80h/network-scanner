import socket

# --- Configuration ---
target_ip = "8.8.8.8"  # Google's Public DNS server, a reliable target
port_to_scan = 53      # Port 53 is for DNS, it should be open on this server

# --- The Scanner Logic ---
try:
    # 1. Create a socket object
    # AF_INET means we are using IPv4 addresses (the standard x.x.x.x format)
    # SOCK_STREAM means we are using TCP (the formal handshake protocol)
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    # Set a timeout so our program doesn't hang forever if a port is unresponsive
    s.settimeout(1) # Wait for a maximum of 1 second

    # 2. Try to connect to the target IP and port
    # connect_ex returns an error indicator. 0 means success.
    result = s.connect_ex((target_ip, port_to_scan))

    # 3. Check the result
    if result == 0:
        print(f"Port {port_to_scan} on {target_ip} is OPEN")
    else:
        print(f"Port {port_to_scan} on {target_ip} is CLOSED")

    # 4. Close the socket to be tidy
    s.close()

except socket.error as e:
    print(f"An error occurred: {e}")