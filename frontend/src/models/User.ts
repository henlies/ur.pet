import { AddressInterface } from "./Address";
import { BloodInterface } from "./Blood";
import { GenderInterface } from "./Gender";
import { PetInterface } from "./Pet";
import { PrefixInterface } from "./Prefix";
import { RoleInterface } from "./Role";

export interface UserInterface {
    ID?: number;
    PrefixID?: number;
    Prefix?: PrefixInterface;
    Firstname: string;
    Lastname: string;
    Nickname: string;
    GenderID?: number;
    Gender?: GenderInterface;
    Phone: string;
    AddressID?: number;
    Address?: AddressInterface;
    Email: string;
    Birth: Date;
    BloodID?: number;
    Blood?: BloodInterface;
    PetID?: number;
    Pet?: PetInterface;
    Descript: string;
    Pic: string;
    User: string;
    Pass: string;
    RoleID?: number;
    Role?: RoleInterface;
    Status: number;
}