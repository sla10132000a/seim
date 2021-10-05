from __future__ import print_function
import logging
import grpc
import image_pb2
import image_pb2_grpc
from concurrent import futures
import predict

class ImageServe(image_pb2_grpc.ImageServiceServicer):

    def ReadImage(self, request, context):
        ml = predict.ml()
        res = ml.predict()
        print(res)
        return image_pb2.Image(content = str(res))

def run():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    image_pb2_grpc.add_ImageServiceServicer_to_server(ImageServe(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig()
    run()