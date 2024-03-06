export interface IUserData {
    ID: number;
    CreatedAt: Date;
    updatedAt: Date | null;
    DeletedAt: Date | null;
    Email: string;
    Username: string;
    UserType: string;
}