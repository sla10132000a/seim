import predict
import grpc
import image_pb2
import image_pb2_grpc

def main():
    ml = predict.ml()
    res = ml.predict()
    print("------------------")
    print(res)
    return image_pb2.Image(content = res)

main()