
## Diagramas de Classe

@startuml

interface SenderFunc {
  + get() : libchan.Sender
}

enum RequestType {
  CreateRequest
  DeleteRequest
  GetRequest
  ListRequest
}

enum ResponseType {
  CreateRequest
  DeleteRequest
  GetRequest
  ListRequest
}

class AnyLittleThing {
  id:   string
  data: string
}

class Request  {
  responseChan: libchan.Sender
  payload:      []byte
  type:         RequestType
}

class AnyLittleThingList {
  + list: List<AnyLittleThing>
}

class Response  {
  payload: []byte
  type:    ResponseType
  err:     error
}

interface AnyLittleThingAdapterIF  {
  listen(libchan.Receiver): error
}

class AnyLittleThingAdapter  {
  thingies: map<string, AnyLittleThing>
  newThingeyAdapter(): ThingeyAdapter

}

enum ClientPeerType {
  Local,
  RemoteByHttp,
  RemoeBySpdy,
  RemoteBySocket,
  RemoteByWebSockect
}

interface ClientPeer {
  getType(): ClientPeerType
  create(obj: object): error
  dispatch(req: Request): Response
  delete(): error
  get(string): object
  list(): List
}

class AnyLittleThingClientPeer {
  type: ClientPeerType
  peerData: object
  senderFunc:   SenderFunc
  receiver:     libchan.Receiver
  remoteSender: libchan.Sender
  create(f: SenderFunc, r: libchan.Receiver, rs: libchan.Sender)
}

@enduml


