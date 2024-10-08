export interface IUser {
  id: number;
  firstName: string;
  lastName: string;
  email: string;
  role: 'operator' | 'admin' | 'user';
}
