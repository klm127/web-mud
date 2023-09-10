--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4 (Debian 15.4-1.pgdg110+1)
-- Dumped by pg_dump version 15.4 (Debian 15.4-1.pgdg110+1)

-- Started on 2023-09-10 02:06:23 UTC

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
-- TOC entry 3338 (class 1262 OID 16385)
-- Name: sdcmud; Type: DATABASE; Schema: -; Owner: sdcadmin
--

CREATE DATABASE sdcmud WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE sdcmud OWNER TO sdcadmin;

\connect sdcmud

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
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 3339 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


--
-- TOC entry 839 (class 1247 OID 16387)
-- Name: userlevel; Type: TYPE; Schema: public; Owner: sdcadmin
--

CREATE TYPE public.userlevel AS ENUM (
    'player',
    'trial-builder',
    'builder',
    'moderator',
    'admin'
);


ALTER TYPE public.userlevel OWNER TO sdcadmin;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 214 (class 1259 OID 16397)
-- Name: rooms; Type: TABLE; Schema: public; Owner: sdcadmin
--

CREATE TABLE public.rooms (
    id bigint NOT NULL
);


ALTER TABLE public.rooms OWNER TO sdcadmin;

--
-- TOC entry 215 (class 1259 OID 16400)
-- Name: users; Type: TABLE; Schema: public; Owner: sdcadmin
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name character varying(20) NOT NULL,
    password character varying NOT NULL,
    created timestamp without time zone DEFAULT now() NOT NULL,
    last_login timestamp without time zone DEFAULT now() NOT NULL,
    level public.userlevel NOT NULL
);


ALTER TABLE public.users OWNER TO sdcadmin;

--
-- TOC entry 216 (class 1259 OID 16407)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: sdcadmin
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO sdcadmin;

--
-- TOC entry 3340 (class 0 OID 0)
-- Dependencies: 216
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sdcadmin
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3183 (class 2604 OID 16408)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: sdcadmin
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3330 (class 0 OID 16397)
-- Dependencies: 214
-- Data for Name: rooms; Type: TABLE DATA; Schema: public; Owner: sdcadmin
--



--
-- TOC entry 3331 (class 0 OID 16400)
-- Dependencies: 215
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: sdcadmin
--



--
-- TOC entry 3341 (class 0 OID 0)
-- Dependencies: 216
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sdcadmin
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- TOC entry 3187 (class 2606 OID 16410)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: sdcadmin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


-- Completed on 2023-09-10 02:06:24 UTC

--
-- PostgreSQL database dump complete
--

