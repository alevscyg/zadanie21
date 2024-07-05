-- CreateTable
CREATE TABLE "User" (
    "id" SERIAL NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "banned" BOOLEAN NOT NULL DEFAULT false,
    "banReason" VARCHAR(100),

    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Role" (
    "value" VARCHAR(30) NOT NULL,
    "description" VARCHAR(255) NOT NULL,

    CONSTRAINT "Role_pkey" PRIMARY KEY ("value")
);

-- CreateTable
CREATE TABLE "UserRole" (
    "userid" INTEGER NOT NULL,
    "value" VARCHAR(30) NOT NULL,

    CONSTRAINT "UserRole_pkey" PRIMARY KEY ("userid","value")
);

-- CreateIndex
CREATE UNIQUE INDEX "User_email_key" ON "User"("email");

-- AddForeignKey
ALTER TABLE "UserRole" ADD CONSTRAINT "UserRole_userid_fkey" FOREIGN KEY ("userid") REFERENCES "User"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "UserRole" ADD CONSTRAINT "UserRole_value_fkey" FOREIGN KEY ("value") REFERENCES "Role"("value") ON DELETE CASCADE ON UPDATE CASCADE;
