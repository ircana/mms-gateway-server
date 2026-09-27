-- +goose Up
-- +goose StatementBegin
ALTER TABLE `messages`
MODIFY COLUMN `type` enum('Text', 'Data', 'Mms') NOT NULL DEFAULT 'Text';
-- +goose StatementEnd
-- +goose StatementBegin
-- An MMS carries its attachments base64-encoded inside the content JSON, which
-- overflows TEXT (64KB) on anything but a thumbnail.
ALTER TABLE `messages`
MODIFY COLUMN `content` longtext NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM `messages` WHERE `type` = 'Mms';
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `messages`
MODIFY COLUMN `type` enum('Text', 'Data') NOT NULL DEFAULT 'Text';
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `messages`
MODIFY COLUMN `content` text NOT NULL;
-- +goose StatementEnd
