import type { NextApiRequest, NextApiResponse } from 'next'
import { genSaltSync, hashSync } from 'bcryptjs';
import jwt from 'jsonwebtoken';
import { setCookie } from 'cookies-next';
import { ISingUp } from '../../../types/user';
import prisma from '../../../../prisma/index'


const handler = async (req: NextApiRequest, res: NextApiResponse) => {

    if (req.method === "POST") {

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

            const refreshToken = jwt.sign({ id: UserData.id }, process.env.SECRET_KEY as string, { expiresIn: fullYear })

            setCookie("token", refreshToken, {
                httpOnly: true,
                secure: process.env.NODE_ENV === "production",
                sameSite: "strict",
                maxAge: fullYear, // full year
                expires: new Date(Date.now() + fullYear),
                path: "/",
                req,
                res
            })

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
}

export default handler;
