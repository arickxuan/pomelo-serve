import socket

def tcp_client(host='127.0.0.1', port=8080):
    """简单的 TCP 客户端"""

    # 创建 socket 对象
    # socket.AF_INET: IPv4, socket.SOCK_STREAM: TCP
    client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    try:
        # 连接到服务器
        print(f"正在连接到 {host}:{port}...")
        client_socket.connect((host, port))
        print("连接成功！")

        # 发送数据
        message = "Hello, Server!"
        print(f"发送: {message}")
        client_socket.send(message.encode('utf-8'))

        # 接收服务器响应
        response = client_socket.recv(1024)  # 1024 是缓冲区大小
        print(f"接收: {response.decode('utf-8')}")

    except ConnectionRefusedError:
        print("连接失败：服务器拒绝连接或未启动")
    except Exception as e:
        print(f"发生错误: {e}")
    finally:
        # 关闭连接
        client_socket.close()
        print("连接已关闭")

if __name__ == "__main__":
    tcp_client()