import { NextRequest, NextResponse } from "next/server";
import { authenticate } from "../../../lib/middleware/authenticate";

export async function GET(req: NextRequest) {
  const authResult = await authenticate(req);
  if (authResult instanceof NextResponse) {
    return authResult; // 認証に失敗した場合はエラーを返す
  }

  // 認証に成功した場合、スケジュールデータを返す処理
  return NextResponse.json({ message: "スケジュールデータ" });
}
