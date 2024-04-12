export interface IUserData {
  ID: number;
  CreatedAt: Date;
  UpdatedAt: Date | null;
  DeletedAt: Date | null;
  Email: string;
  Username: string;
  UserType: string;
}
