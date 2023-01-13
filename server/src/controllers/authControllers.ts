import prisma from '../../prisma/index';
import jwt from 'jsonwebtoken';
import { compareSync, genSaltSync, hashSync } from 'bcryptjs';
import { Request, Response } from 'express';
import { ILogin, ISingUp } from '../types/user';

export  const SingUp = async (req: Request, res: Response) => {
        try {

            const { password, email, firstName, lastName }: ISingUp = req.body;

            if (!lastName || !password || !email || !firstName) return res.status(400).json({ massage: "InValid Data" });
            if (!(password.length > 6) && !(lastName.length > 3) && !(firstName.length > 3) && !(email.length > 8))
                return res.status(400).json({ massage: 'unValid Fields' });

            const user = await prisma.user.findFirst({ where: { email: email }, select: { email: true } });

            if (user?.email) return res.status(400).json({ error: "user already exist try login", user })

            const salt = genSaltSync(10);
            const hashPassword = hashSync(password, salt);

            const UserData = await prisma.user.create({
                data: {
                    firstName: firstName,
                    lastName: lastName,
                    email: email,
                    password: hashPassword
                }
            })

            const fullYear = 1000 * 60 * 60 * 24 * 365;

            const token = jwt.sign({ id: UserData.id }, process.env.SECRET_KEY as string, { expiresIn: fullYear })

            res.setHeader('Set-Cookie', `token=${token}; Same-Site=Strict; HttpOnly=True; Path=/; Expires=${new Date(Date.now() + fullYear)} Max-Age=${fullYear}; Secure=${process.env.NODE_ENV === "production" ? "True" : "False"}`);

            const data = {
                id: UserData.id,
                createdAt: UserData.createdAt,
                email: UserData.email,
                lastName: UserData.lastName,
                firstName: UserData.firstName,
            }

            return res.status(200).json({ data, massage: "sing up success" })

        } catch (error) {
            console.log(error)
            return res.status(500).json({ massage: "internal Server Error" })
        }
}



export const Login = async (req: Request, res: Response) => {

        try {
            const { email, password } = req.body as ILogin;
            const UserData = await prisma.user.findUnique({ where: { email: email } });

            if (!UserData) return res.status(400).json({ error: `user with this email ${email} dose not exist` })

            if (!(password.length > 6) && !UserData.password && !(email.length > 8))
            return res.status(400).json({ error: 'unValid Fields' });

            const isMatch = compareSync(password, UserData.password)

            if (!isMatch) return res.status(400).json({ error: `password is incorrect` })

            const fullYear = 1000 * 60 * 60 * 24 * 365;

            const token = jwt.sign({ id: UserData.id }, process.env.SECRET_KEY as string, { expiresIn: fullYear })

            res.setHeader('Set-Cookie', `token=${token}; Same-Site=Strict; HttpOnly=True; Path=/; Expires=${new Date(Date.now() + fullYear)} Max-Age=${fullYear}; Secure=${process.env.NODE_ENV === "production" ? "True" : "False"}`);

            const data = {
                id: UserData.id,
                createdAt: UserData.createdAt,
                email: UserData.email,
                lastName: UserData.lastName,
                firstName: UserData.firstName
            }

            return res.status(200).json({ data, massage: "login success" })

        } catch (error) {
            console.log(error)
            return res.status(500).json({ massage: "internal Server Error" })
        }
}

export const Logout = async (req: Request, res: Response) => {
    try {
        res.clearCookie("token");
        res.status(200).json({ message: 'Logout success' });
    } catch (error) {
        console.log(error)
        return res.status(500).json({ massage: "internal Server Error" })
    }
}


