import path from 'path'
import grpc from 'grpc'
import protoLoader from '@grpc/proto-loader'

const __dirname = path.resolve()

const ProtoFile = `${__dirname}/../service.proto`

const packageDefinition = protoLoader.loadSync(
  ProtoFile,
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
  }
)

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition)

const MyService = protoDescriptor.test.MyService

function doSomething(call, callback) {
  console.info('doSomething called', call.request)
  callback(null, {
    message: `This is the response from the GRPC service for ${call.request.name}`
  })
}

const server = new grpc.Server()
server.addService(MyService.service, {
  DoSomething: doSomething
})

console.info('Iniciando servidor MyService em 50051')
server.bind('0.0.0.0:50051', grpc.ServerCredentials.createInsecure())
server.start()
