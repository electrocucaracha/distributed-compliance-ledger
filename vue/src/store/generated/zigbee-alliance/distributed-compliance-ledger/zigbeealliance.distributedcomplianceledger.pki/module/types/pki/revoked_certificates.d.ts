import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";
export interface RevokedCertificates {
    subject: string;
    subjectKeyId: string;
    pemCert: string;
    serialNumber: string;
    issuer: string;
    authorityKeyId: string;
    rootSubject: string;
    rootSubjectKeyId: string;
    isRoot: boolean;
    owner: string;
}
export declare const RevokedCertificates: {
    encode(message: RevokedCertificates, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): RevokedCertificates;
    fromJSON(object: any): RevokedCertificates;
    toJSON(message: RevokedCertificates): unknown;
    fromPartial(object: DeepPartial<RevokedCertificates>): RevokedCertificates;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
