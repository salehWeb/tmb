import prisma from '../../../../prisma/index';
import jwt from 'jsonwebtoken';
import { compareSync } from 'bcryptjs';
import { NextApiRequest, NextApiResponse } from 'next';
import { setCookie } from 'cookies-next';
import { ILogin } from '../../../types/user';


const handler = async (req: NextApiRequest, res: NextApiResponse) => {

    if (req.method === "POST") {

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

            setCookie("token", token, {
                httpOnly: true,
                secure: process.env.NODE_ENV === "production",
                sameSite: "strict",
                maxAge: fullYear, // full year
                expires: new Date(Date.now() + fullYear),
                path: "/",
                req,
                res
            });

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

    };
}

export default handler;
