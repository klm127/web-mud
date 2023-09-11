--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4 (Debian 15.4-1.pgdg110+1)
-- Dumped by pg_dump version 15.4 (Debian 15.4-1.pgdg110+1)

-- Started on 2023-09-11 17:38:24 UTC

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE IF EXISTS sdcmud;
--
-- TOC entry 3357 (class 1262 OID 16385)
-- Name: sdcmud; Type: DATABASE; Schema: -; Owner: sdcadmin
--

CREATE DATABASE sdcmud WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE sdcmud OWNER TO sdcadmin;


SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 6 (class 2615 OID 16386)
-- Name: mud; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA mud;


ALTER SCHEMA mud OWNER TO pg_database_owner;

--
-- TOC entry 3358 (class 0 OID 0)
-- Dependencies: 6
-- Name: SCHEMA mud; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA mud IS 'standard public schema';


--
-- TOC entry 843 (class 1247 OID 16388)
-- Name: userlevel; Type: TYPE; Schema: mud; Owner: sdcadmin
--

CREATE TYPE mud.userlevel AS ENUM (
    'player',
    'trial-builder',
    'builder',
    'moderator',
    'admin'
);


ALTER TYPE mud.userlevel OWNER TO sdcadmin;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 220 (class 1259 OID 16427)
-- Name: beings; Type: TABLE; Schema: mud; Owner: sdcadmin
--

CREATE TABLE mud.beings (
    id bigint NOT NULL,
    name character varying NOT NULL,
    description character varying NOT NULL,
    room bigint NOT NULL,
    owner bigint
);


ALTER TABLE mud.beings OWNER TO sdcadmin;

--
-- TOC entry 219 (class 1259 OID 16426)
-- Name: beings_id_seq; Type: SEQUENCE; Schema: mud; Owner: sdcadmin
--

CREATE SEQUENCE mud.beings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE mud.beings_id_seq OWNER TO sdcadmin;

--
-- TOC entry 3359 (class 0 OID 0)
-- Dependencies: 219
-- Name: beings_id_seq; Type: SEQUENCE OWNED BY; Schema: mud; Owner: sdcadmin
--

ALTER SEQUENCE mud.beings_id_seq OWNED BY mud.beings.id;


--
-- TOC entry 218 (class 1259 OID 16418)
-- Name: rooms; Type: TABLE; Schema: mud; Owner: sdcadmin
--

CREATE TABLE mud.rooms (
    id bigint NOT NULL,
    name character varying NOT NULL,
    description character varying NOT NULL,
    img character varying,
    objects bigint[] NOT NULL,
    n bigint,
    s bigint,
    e bigint,
    w bigint,
    ne bigint,
    se bigint,
    sw bigint,
    nw bigint,
    u bigint,
    d bigint
);


ALTER TABLE mud.rooms OWNER TO sdcadmin;

--
-- TOC entry 217 (class 1259 OID 16417)
-- Name: rooms_id_seq; Type: SEQUENCE; Schema: mud; Owner: sdcadmin
--

CREATE SEQUENCE mud.rooms_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE mud.rooms_id_seq OWNER TO sdcadmin;

--
-- TOC entry 3360 (class 0 OID 0)
-- Dependencies: 217
-- Name: rooms_id_seq; Type: SEQUENCE OWNED BY; Schema: mud; Owner: sdcadmin
--

ALTER SEQUENCE mud.rooms_id_seq OWNED BY mud.rooms.id;


--
-- TOC entry 215 (class 1259 OID 16402)
-- Name: users; Type: TABLE; Schema: mud; Owner: sdcadmin
--

CREATE TABLE mud.users (
    id bigint NOT NULL,
    name character varying(20) NOT NULL,
    password character varying NOT NULL,
    created timestamp without time zone DEFAULT now() NOT NULL,
    last_login timestamp without time zone DEFAULT now() NOT NULL,
    level mud.userlevel NOT NULL,
    being bigint NOT NULL
);


ALTER TABLE mud.users OWNER TO sdcadmin;

--
-- TOC entry 216 (class 1259 OID 16409)
-- Name: users_id_seq; Type: SEQUENCE; Schema: mud; Owner: sdcadmin
--

CREATE SEQUENCE mud.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE mud.users_id_seq OWNER TO sdcadmin;

--
-- TOC entry 3361 (class 0 OID 0)
-- Dependencies: 216
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: mud; Owner: sdcadmin
--

ALTER SEQUENCE mud.users_id_seq OWNED BY mud.users.id;


--
-- TOC entry 3194 (class 2604 OID 16430)
-- Name: beings id; Type: DEFAULT; Schema: mud; Owner: sdcadmin
--

ALTER TABLE ONLY mud.beings ALTER COLUMN id SET DEFAULT nextval('mud.beings_id_seq'::regclass);


--
-- TOC entry 3193 (class 2604 OID 16421)
-- Name: rooms id; Type: DEFAULT; Schema: mud; Owner: sdcadmin
--

ALTER TABLE ONLY mud.rooms ALTER COLUMN id SET DEFAULT nextval('mud.rooms_id_seq'::regclass);


--
-- TOC entry 3190 (class 2604 OID 16410)
-- Name: users id; Type: DEFAULT; Schema: mud; Owner: sdcadmin
--

ALTER TABLE ONLY mud.users ALTER COLUMN id SET DEFAULT nextval('mud.users_id_seq'::regclass);


--
-- TOC entry 3351 (class 0 OID 16427)
-- Dependencies: 220
-- Data for Name: beings; Type: TABLE DATA; Schema: mud; Owner: sdcadmin
--



--
-- TOC entry 3349 (class 0 OID 16418)
-- Dependencies: 218
-- Data for Name: rooms; Type: TABLE DATA; Schema: mud; Owner: sdcadmin
--



--
-- TOC entry 3346 (class 0 OID 16402)
-- Dependencies: 215
-- Data for Name: users; Type: TABLE DATA; Schema: mud; Owner: sdcadmin
--



--
-- TOC entry 3362 (class 0 OID 0)
-- Dependencies: 219
-- Name: beings_id_seq; Type: SEQUENCE SET; Schema: mud; Owner: sdcadmin
--

SELECT pg_catalog.setval('mud.beings_id_seq', 1, false);


--
-- TOC entry 3363 (class 0 OID 0)
-- Dependencies: 217
-- Name: rooms_id_seq; Type: SEQUENCE SET; Schema: mud; Owner: sdcadmin
--

SELECT pg_catalog.setval('mud.rooms_id_seq', 1, false);


--
-- TOC entry 3364 (class 0 OID 0)
-- Dependencies: 216
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: mud; Owner: sdcadmin
--

SELECT pg_catalog.setval('mud.users_id_seq', 2, true);


--
-- TOC entry 3200 (class 2606 OID 16434)
-- Name: beings beings_pkey; Type: CONSTRAINT; Schema: mud; Owner: sdcadmin
--

ALTER TABLE ONLY mud.beings
    ADD CONSTRAINT beings_pkey PRIMARY KEY (id);


--
-- TOC entry 3198 (class 2606 OID 16425)
-- Name: rooms rooms_pkey; Type: CONSTRAINT; Schema: mud; Owner: sdcadmin
--

ALTER TABLE ONLY mud.rooms
    ADD CONSTRAINT rooms_pkey PRIMARY KEY (id);


--
-- TOC entry 3202 (class 2606 OID 16436)
-- Name: beings unique_name; Type: CONSTRAINT; Schema: mud; Owner: sdcadmin
--

ALTER TABLE ONLY mud.beings
    ADD CONSTRAINT unique_name UNIQUE (name);


--
-- TOC entry 3196 (class 2606 OID 16412)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: mud; Owner: sdcadmin
--

ALTER TABLE ONLY mud.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3203 (class 2606 OID 16437)
-- Name: users being_fk; Type: FK CONSTRAINT; Schema: mud; Owner: sdcadmin
--

ALTER TABLE ONLY mud.users
    ADD CONSTRAINT being_fk FOREIGN KEY (being) REFERENCES mud.beings(id) NOT VALID;


-- Completed on 2023-09-11 17:38:24 UTC

--
-- PostgreSQL database dump complete
--

