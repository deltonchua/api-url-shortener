-- +goose Up 
CREATE TABLE "url" (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    public_id CHAR(8) NOT NULL,
    url VARCHAR(255) NOT NULL,
    count BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6)
);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP(6);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON "url"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

-- +goose Down
DROP TRIGGER IF EXISTS trigger_update_updated_at ON "url";

DROP FUNCTION IF EXISTS update_updated_at();

DROP TABLE IF EXISTS "url";
