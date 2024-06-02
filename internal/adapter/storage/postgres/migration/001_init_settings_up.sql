CREATE OR REPLACE FUNCTION generate_generic_id()
RETURNS TRIGGER AS $$
DECLARE
    date_part VARCHAR(8);
    seq_id INT;
    seq_name TEXT;
BEGIN
    date_part := TO_CHAR(NEW.created_at, 'YYYYMMDD');
    seq_name := TG_TABLE_NAME || '_id_seq';
    EXECUTE 'SELECT NEXTVAL(''' || seq_name || ''')' INTO seq_id;
    NEW.id := date_part || seq_id::TEXT;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
DO $$
DECLARE
    table_record RECORD;
    sequence_name TEXT;
    trigger_sql TEXT;
    sequence_sql TEXT;
BEGIN
    FOR table_record IN
        SELECT tc.table_name
        FROM information_schema.table_constraints tc
        JOIN information_schema.constraint_column_usage AS ccu
        USING (constraint_schema, constraint_name)
        WHERE tc.constraint_type = 'PRIMARY KEY'
        AND ccu.column_name = 'id'
        AND tc.table_schema = 'public'
    LOOP
        sequence_name := table_record.table_name || '_id_seq';
        sequence_sql := 'CREATE SEQUENCE IF NOT EXISTS ' || quote_ident(sequence_name) || ';';
        EXECUTE sequence_sql;
        trigger_sql := 'CREATE TRIGGER before_insert_' || quote_ident(table_record.table_name) ||
                       ' BEFORE INSERT ON ' || quote_ident(table_record.table_name) ||
                       ' FOR EACH ROW EXECUTE FUNCTION generate_generic_id();';
        EXECUTE trigger_sql;
    END LOOP;
END $$;