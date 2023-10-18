import { InferModel, sql } from "drizzle-orm";
import { integer, sqliteTable, text } from "drizzle-orm/sqlite-core";

export const users = sqliteTable("songs", {
  id: integer("id").primaryKey(),
  firstName: text("title"),
  createdAt: text("created_at").default(sql`CURRENT_TIMESTAMP`),
});

export type User = InferModel<typeof users>;
export type InsertUser = InferModel<typeof users, "insert">;
