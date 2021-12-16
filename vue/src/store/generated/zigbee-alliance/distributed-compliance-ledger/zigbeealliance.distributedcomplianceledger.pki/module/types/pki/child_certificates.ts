/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface ChildCertificates {
  issuer: string
  authorityKeyId: string
  subject: string
  subjectKeyId: string
}

const baseChildCertificates: object = { issuer: '', authorityKeyId: '', subject: '', subjectKeyId: '' }

export const ChildCertificates = {
  encode(message: ChildCertificates, writer: Writer = Writer.create()): Writer {
    if (message.issuer !== '') {
      writer.uint32(10).string(message.issuer)
    }
    if (message.authorityKeyId !== '') {
      writer.uint32(18).string(message.authorityKeyId)
    }
    if (message.subject !== '') {
      writer.uint32(26).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(34).string(message.subjectKeyId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): ChildCertificates {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseChildCertificates } as ChildCertificates
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.issuer = reader.string()
          break
        case 2:
          message.authorityKeyId = reader.string()
          break
        case 3:
          message.subject = reader.string()
          break
        case 4:
          message.subjectKeyId = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): ChildCertificates {
    const message = { ...baseChildCertificates } as ChildCertificates
    if (object.issuer !== undefined && object.issuer !== null) {
      message.issuer = String(object.issuer)
    } else {
      message.issuer = ''
    }
    if (object.authorityKeyId !== undefined && object.authorityKeyId !== null) {
      message.authorityKeyId = String(object.authorityKeyId)
    } else {
      message.authorityKeyId = ''
    }
    if (object.subject !== undefined && object.subject !== null) {
      message.subject = String(object.subject)
    } else {
      message.subject = ''
    }
    if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
      message.subjectKeyId = String(object.subjectKeyId)
    } else {
      message.subjectKeyId = ''
    }
    return message
  },

  toJSON(message: ChildCertificates): unknown {
    const obj: any = {}
    message.issuer !== undefined && (obj.issuer = message.issuer)
    message.authorityKeyId !== undefined && (obj.authorityKeyId = message.authorityKeyId)
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    return obj
  },

  fromPartial(object: DeepPartial<ChildCertificates>): ChildCertificates {
    const message = { ...baseChildCertificates } as ChildCertificates
    if (object.issuer !== undefined && object.issuer !== null) {
      message.issuer = object.issuer
    } else {
      message.issuer = ''
    }
    if (object.authorityKeyId !== undefined && object.authorityKeyId !== null) {
      message.authorityKeyId = object.authorityKeyId
    } else {
      message.authorityKeyId = ''
    }
    if (object.subject !== undefined && object.subject !== null) {
      message.subject = object.subject
    } else {
      message.subject = ''
    }
    if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
      message.subjectKeyId = object.subjectKeyId
    } else {
      message.subjectKeyId = ''
    }
    return message
  }
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>
