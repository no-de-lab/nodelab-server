CREATE TRIGGER create_user
AFTER INSERT ON account
FOR EACH ROW
    INSERT INTO user (email, created_at)
    VALUES (NEW.email, NOW());
