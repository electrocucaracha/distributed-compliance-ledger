import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";
export interface ApprovedCertificates {
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
export declare const ApprovedCertificates: {
    encode(message: ApprovedCertificates, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ApprovedCertificates;
    fromJSON(object: any): ApprovedCertificates;
    toJSON(message: ApprovedCertificates): unknown;
    fromPartial(object: DeepPartial<ApprovedCertificates>): ApprovedCertificates;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
