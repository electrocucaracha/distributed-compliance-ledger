/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface RevokedCertificates {
  subject: string
  subjectKeyId: string
  pemCert: string
  serialNumber: string
  issuer: string
  authorityKeyId: string
  rootSubject: string
  rootSubjectKeyId: string
  isRoot: boolean
  owner: string
}

const baseRevokedCertificates: object = {
  subject: '',
  subjectKeyId: '',
  pemCert: '',
  serialNumber: '',
  issuer: '',
  authorityKeyId: '',
  rootSubject: '',
  rootSubjectKeyId: '',
  isRoot: false,
  owner: ''
}

export const RevokedCertificates = {
  encode(message: RevokedCertificates, writer: Writer = Writer.create()): Writer {
    if (message.subject !== '') {
      writer.uint32(10).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(18).string(message.subjectKeyId)
    }
    if (message.pemCert !== '') {
      writer.uint32(26).string(message.pemCert)
    }
    if (message.serialNumber !== '') {
      writer.uint32(34).string(message.serialNumber)
    }
    if (message.issuer !== '') {
      writer.uint32(42).string(message.issuer)
    }
    if (message.authorityKeyId !== '') {
      writer.uint32(50).string(message.authorityKeyId)
    }
    if (message.rootSubject !== '') {
      writer.uint32(58).string(message.rootSubject)
    }
    if (message.rootSubjectKeyId !== '') {
      writer.uint32(66).string(message.rootSubjectKeyId)
    }
    if (message.isRoot === true) {
      writer.uint32(72).bool(message.isRoot)
    }
    if (message.owner !== '') {
      writer.uint32(82).string(message.owner)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RevokedCertificates {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRevokedCertificates } as RevokedCertificates
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string()
          break
        case 2:
          message.subjectKeyId = reader.string()
          break
        case 3:
          message.pemCert = reader.string()
          break
        case 4:
          message.serialNumber = reader.string()
          break
        case 5:
          message.issuer = reader.string()
          break
        case 6:
          message.authorityKeyId = reader.string()
          break
        case 7:
          message.rootSubject = reader.string()
          break
        case 8:
          message.rootSubjectKeyId = reader.string()
          break
        case 9:
          message.isRoot = reader.bool()
          break
        case 10:
          message.owner = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RevokedCertificates {
    const message = { ...baseRevokedCertificates } as RevokedCertificates
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
    if (object.pemCert !== undefined && object.pemCert !== null) {
      message.pemCert = String(object.pemCert)
    } else {
      message.pemCert = ''
    }
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      message.serialNumber = String(object.serialNumber)
    } else {
      message.serialNumber = ''
    }
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
    if (object.rootSubject !== undefined && object.rootSubject !== null) {
      message.rootSubject = String(object.rootSubject)
    } else {
      message.rootSubject = ''
    }
    if (object.rootSubjectKeyId !== undefined && object.rootSubjectKeyId !== null) {
      message.rootSubjectKeyId = String(object.rootSubjectKeyId)
    } else {
      message.rootSubjectKeyId = ''
    }
    if (object.isRoot !== undefined && object.isRoot !== null) {
      message.isRoot = Boolean(object.isRoot)
    } else {
      message.isRoot = false
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner)
    } else {
      message.owner = ''
    }
    return message
  },

  toJSON(message: RevokedCertificates): unknown {
    const obj: any = {}
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    message.pemCert !== undefined && (obj.pemCert = message.pemCert)
    message.serialNumber !== undefined && (obj.serialNumber = message.serialNumber)
    message.issuer !== undefined && (obj.issuer = message.issuer)
    message.authorityKeyId !== undefined && (obj.authorityKeyId = message.authorityKeyId)
    message.rootSubject !== undefined && (obj.rootSubject = message.rootSubject)
    message.rootSubjectKeyId !== undefined && (obj.rootSubjectKeyId = message.rootSubjectKeyId)
    message.isRoot !== undefined && (obj.isRoot = message.isRoot)
    message.owner !== undefined && (obj.owner = message.owner)
    return obj
  },

  fromPartial(object: DeepPartial<RevokedCertificates>): RevokedCertificates {
    const message = { ...baseRevokedCertificates } as RevokedCertificates
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
    if (object.pemCert !== undefined && object.pemCert !== null) {
      message.pemCert = object.pemCert
    } else {
      message.pemCert = ''
    }
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      message.serialNumber = object.serialNumber
    } else {
      message.serialNumber = ''
    }
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
    if (object.rootSubject !== undefined && object.rootSubject !== null) {
      message.rootSubject = object.rootSubject
    } else {
      message.rootSubject = ''
    }
    if (object.rootSubjectKeyId !== undefined && object.rootSubjectKeyId !== null) {
      message.rootSubjectKeyId = object.rootSubjectKeyId
    } else {
      message.rootSubjectKeyId = ''
    }
    if (object.isRoot !== undefined && object.isRoot !== null) {
      message.isRoot = object.isRoot
    } else {
      message.isRoot = false
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner
    } else {
      message.owner = ''
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