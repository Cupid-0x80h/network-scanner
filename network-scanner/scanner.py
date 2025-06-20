import socket
import threading
from queue import Queue

# The worker function remains exactly the same
def scan_port(target_ip, port_queue, open_ports_list):
    while not port_queue.empty():
        port = port_queue.get()
        try:
            s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            socket.setdefaulttimeout(0.5)
            result = s.connect_ex((target_ip, port))
            if result == 0:
                open_ports_list.append(port)
            s.close()
        except socket.error:
            pass
        port_queue.task_done()

# --- The Main Program (The Manager) ---

# 1. GET USER INPUT
target = input("Enter the target IP address to scan: ")
print("\nSelect scan type:")
print("1. Default Ports (1-1024)")
print("2. All Ports (1-65535)")
choice = input("Enter your choice (1 or 2): ")

start_port = 0
end_port = 0

if choice == '1':
    start_port = 1
    end_port = 1024
elif choice == '2':
    start_port = 1
    end_port = 65535
else:
    print("Invalid choice. Exiting.")
    exit()

print("-" * 50)
print(f"Scanning {target} for open ports...")
print("-" * 50)


# 2. SETUP THE QUEUE AND WORKERS
port_q = Queue()
open_ports = []

# Fill the queue with the chosen port range
for port in range(start_port, end_port + 1):
    port_q.put(port)

# Create and start 100 threads
threads = []
for _ in range(100):
    # Pass arguments to the worker function
    thread = threading.Thread(target=scan_port, args=(target, port_q, open_ports))
    thread.start()
    threads.append(thread)

# Wait for the queue to be empty
port_q.join()

# 3. DISPLAY RESULTS
open_ports.sort()
if open_ports:
    print("Open ports found:")
    for port in open_ports:
        print(port)
else:
    print("No open ports found in the specified range.")
