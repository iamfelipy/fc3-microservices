-- CreateTable
CREATE TABLE "accounts" (
    "id" TEXT NOT NULL,
    "client_id" TEXT NOT NULL,
    "balance" DECIMAL(65,30) NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "accounts_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "accounts_client_id_key" ON "accounts"("client_id");
