import { NextRequest, NextResponse } from "next/server";
import jwt from "jsonwebtoken";

export async function authenticate(req: NextRequest) {
  const authHeader = req.headers.get("authorization");
  if (!authHeader) {
    return NextResponse.json(
      { message: "user not authenticated" },
      { status: 401 }
    );
  }

  const token = authHeader.split(" ")[1];
  if (!token) {
    return NextResponse.json(
      { message: "user not authenticated" },
      { status: 401 }
    );
  }

  try {
    const decoded = jwt.verify(token, process.env.JWT_SECRET as string);
    req.headers.set("user", JSON.stringify(decoded)); // デコードされたユーザー情報をリクエストに追加
  } catch (error) {
    return NextResponse.json(
      { message: "token invalid or expired" },
      { status: 403 }
    );
  }
}
