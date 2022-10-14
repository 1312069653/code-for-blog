import grpc
import stringdb_pb2
import stringdb_pb2_grpc

PORT = 4050
TIMEOUT_SECONDS = 3


def set_value(stub, key, value):
    request = stringdb_pb2.SetValueRequest(key=key, value=value)
    response = stub.SetValue(request, TIMEOUT_SECONDS)
    return response.value

def get_value(stub, key):
    request = stringdb_pb2.GetValueRequest(key=key)
    response = stub.GetValue(request, TIMEOUT_SECONDS)
    return response.value

def count_value(stub, key):
    request = stringdb_pb2.CountValueRequest(key=key)
    response = stub.CountValue(request, TIMEOUT_SECONDS)
    return response.count


def main():
  channel = grpc.insecure_channel('localhost:4050')
  stub = stringdb_pb2_grpc.StringDbStub(channel)

  set_value(stub, 'foo', 'bar')
  set_value(stub, 'baz', 'anaconda is here')
  print(get_value(stub, 'foo'))
  print(count_value(stub, 'baz'))


if __name__ == '__main__':
  main()
